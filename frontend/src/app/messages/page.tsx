'use client'
import { useState, useEffect, useRef } from 'react'
import { useRouter } from 'next/navigation'
import { fetchMe, getUser } from '@/lib/auth'
import api, { ApiError } from '@/lib/api'

interface User {
  id: number
  username: string
  email: string
  display_name: string
}

interface Message {
  id: number
  sender_id: number
  receiver_id: number
  content: string
  created_at: string
}

export default function MessagesPage() {
  const router = useRouter()
  const [users, setUsers] = useState<User[]>([])
  const [selectedUser, setSelectedUser] = useState<User | null>(null)
  const [messages, setMessages] = useState<Message[]>([])
  const [newMessage, setNewMessage] = useState('')
  const [sending, setSending] = useState(false)
  const [loadingUsers, setLoadingUsers] = useState(true)
  const [authChecking, setAuthChecking] = useState(true)
  const [error, setError] = useState('')
  const [currentUser, setCurrentUser] = useState<User | null>(null)
  const messagesEndRef = useRef<HTMLDivElement>(null)
  const pollRef = useRef<ReturnType<typeof setInterval> | null>(null)

  // Auth check: use sessionStorage cache first (fast), fall back to /auth/me via httpOnly cookie
  useEffect(() => {
    const cached = getUser()
    if (cached) {
      setCurrentUser(cached)
      setAuthChecking(false)
      return
    }
    // No cache - verify via backend cookie (handles new tab / hard refresh)
    ;(async () => {
      const user = await fetchMe()
      if (!user) {
        router.push('/login')
      } else {
        setCurrentUser(user)
        setAuthChecking(false)
      }
    })()
  }, [router])

  // Fetch user list - runs after auth confirmed
  async function loadUsers() {
    const me = getUser()
    try {
      const data = await api.get<{ data: User[] }>('/users')
      const others = data.data.filter((u: User) => u.id !== me?.id)
      setUsers(others)
      setError('')
    } catch (err) {
      setUsers(prev => {
        if (prev.length === 0) setError(err instanceof ApiError ? err.message : 'Khong the tai danh sach nguoi dung')
        return prev
      })
    } finally {
      setLoadingUsers(false)
    }
  }

  useEffect(() => {
    if (!authChecking) loadUsers()
  }, [authChecking])

  async function loadMessages(userId: number) {
    try {
      const data = await api.get<{ data: Message[] }>(`/messages/${userId}`)
      setMessages(data.data)
    } catch {
      // silent poll failure - don't disrupt UX
    }
  }

  useEffect(() => {
    if (!selectedUser) return
    loadMessages(selectedUser.id)
    pollRef.current = setInterval(() => loadMessages(selectedUser.id), 2000)
    return () => {
      if (pollRef.current) clearInterval(pollRef.current)
    }
  }, [selectedUser])

  async function handleSend(e: React.FormEvent) {
    e.preventDefault()
    if (!newMessage.trim() || !selectedUser) return
    setSending(true)
    try {
      await api.post('/messages', { receiver_id: selectedUser.id, content: newMessage.trim() })
      setNewMessage('')
      await loadMessages(selectedUser.id)
    } catch (err) {
      setError(err instanceof ApiError ? err.message : 'Khong the gui tin nhan')
    } finally {
      setSending(false)
    }
  }

  function selectUser(user: User) {
    setSelectedUser(user)
    setMessages([])
  }

  if (authChecking) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="text-gray-400 text-sm">Dang kiem tra dang nhap...</div>
      </div>
    )
  }

  return (
    <div className="max-w-6xl mx-auto px-4 py-8">
      <div className="bg-gradient-to-br from-blue-700 to-indigo-800 rounded-2xl p-6 text-white mb-6 shadow-lg">
        <h1 className="text-2xl font-bold flex items-center gap-2">
          <span>💬</span> Tin nhan
        </h1>
        <p className="text-blue-200 text-sm mt-1">Nhan tin truc tiep voi ban hoc</p>
      </div>

      {error && (
        <div className="mb-4 p-3 bg-red-50 border border-red-200 text-red-700 rounded-lg text-sm">{error}</div>
      )}

      <div className="bg-white rounded-2xl shadow-md overflow-hidden flex" style={{ height: '600px' }}>
        {/* User list */}
        <div className="w-72 border-r border-gray-200 flex flex-col">
          <div className="p-4 border-b border-gray-100 bg-gray-50">
            <h2 className="font-semibold text-gray-700 text-sm uppercase tracking-wide">
              Nguoi dung ({users.length})
            </h2>
          </div>
          <div className="overflow-y-auto flex-1">
            {loadingUsers ? (
              <div className="p-4 text-center text-gray-400 text-sm">Dang tai...</div>
            ) : users.length === 0 ? (
              <div className="p-4 text-center text-gray-400 text-sm">Khong co nguoi dung nao</div>
            ) : (
              users.map(user => (
                <button
                  key={user.id}
                  onClick={() => selectUser(user)}
                  className={`w-full text-left px-4 py-3 flex items-center gap-3 hover:bg-blue-50 transition-colors border-b border-gray-50 ${
                    selectedUser?.id === user.id ? 'bg-blue-50 border-l-4 border-l-blue-600' : ''
                  }`}
                >
                  <div className="w-10 h-10 rounded-full bg-gradient-to-br from-blue-400 to-indigo-500 flex items-center justify-center text-white font-bold text-sm flex-shrink-0">
                    {(user.display_name || user.username).charAt(0).toUpperCase()}
                  </div>
                  <div className="min-w-0">
                    <div className="font-medium text-gray-800 text-sm truncate">{user.display_name || user.username}</div>
                    <div className="text-xs text-gray-400 truncate">@{user.username}</div>
                  </div>
                </button>
              ))
            )}
          </div>
        </div>

        {/* Conversation panel */}
        <div className="flex-1 flex flex-col">
          {selectedUser ? (
            <>
              <div className="px-6 py-4 border-b border-gray-200 bg-gray-50 flex items-center gap-3">
                <div className="w-9 h-9 rounded-full bg-gradient-to-br from-blue-400 to-indigo-500 flex items-center justify-center text-white font-bold text-sm">
                  {(selectedUser.display_name || selectedUser.username).charAt(0).toUpperCase()}
                </div>
                <div>
                  <div className="font-semibold text-gray-800">{selectedUser.display_name || selectedUser.username}</div>
                  <div className="text-xs text-gray-400">@{selectedUser.username}</div>
                </div>
              </div>

              <div className="flex-1 overflow-y-auto p-4 space-y-3">
                {messages.length === 0 ? (
                  <div className="text-center text-gray-400 text-sm mt-8">
                    Chua co tin nhan nao. Hay bat dau cuoc tro chuyen!
                  </div>
                ) : (
                  messages.map(msg => {
                    const isMine = msg.sender_id === currentUser?.id
                    return (
                      <div key={msg.id} className={`flex ${isMine ? 'justify-end' : 'justify-start'}`}>
                        <div className={`max-w-xs lg:max-w-md px-4 py-2 rounded-2xl text-sm ${
                          isMine
                            ? 'bg-gradient-to-r from-blue-600 to-indigo-700 text-white rounded-br-sm'
                            : 'bg-gray-100 text-gray-800 rounded-bl-sm'
                        }`}>
                          <p>{msg.content}</p>
                          <p className={`text-xs mt-1 ${isMine ? 'text-blue-200' : 'text-gray-400'}`}>
                            {new Date(msg.created_at).toLocaleTimeString('vi-VN', { hour: '2-digit', minute: '2-digit' })}
                          </p>
                        </div>
                      </div>
                    )
                  })
                )}
                <div ref={messagesEndRef} />
              </div>

              <form onSubmit={handleSend} className="p-4 border-t border-gray-200 flex gap-3">
                <input
                  type="text"
                  value={newMessage}
                  onChange={e => setNewMessage(e.target.value)}
                  placeholder={`Nhan tin cho ${selectedUser.display_name || selectedUser.username}...`}
                  className="flex-1 px-4 py-2.5 border border-gray-300 rounded-full focus:outline-none focus:ring-2 focus:ring-blue-500 text-gray-900 text-sm"
                />
                <button
                  type="submit"
                  disabled={sending || !newMessage.trim()}
                  className="px-5 py-2.5 bg-gradient-to-r from-blue-600 to-indigo-700 text-white rounded-full font-medium text-sm hover:opacity-90 transition disabled:opacity-50 flex items-center gap-1"
                >
                  <span>📤</span><span>Gui</span>
                </button>
              </form>
            </>
          ) : (
            <div className="flex-1 flex items-center justify-center text-gray-400">
              <div className="text-center">
                <div className="text-5xl mb-3">💬</div>
                <p className="font-medium">Chon nguoi dung de bat dau nhan tin</p>
                <p className="text-sm mt-1">Danh sach ben trai hien thi tat ca nguoi dung</p>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  )
}
