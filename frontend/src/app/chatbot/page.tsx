'use client'
import { useState, useRef, useEffect } from 'react'
import api, { ApiError } from '@/lib/api'

interface ChatMessage {
  id: number; session_id: string; role: 'user' | 'assistant'
  content: string; created_at: string
}

function genId() { return 'sess_' + Math.random().toString(36).substr(2, 9) + '_' + Date.now() }

export default function ChatbotPage() {
  const [messages, setMessages] = useState<ChatMessage[]>([])
  const [input, setInput] = useState('')
  const [loading, setLoading] = useState(false)
  const [sessionId] = useState(() => {
    if (typeof window !== 'undefined') {
      const s = sessionStorage.getItem('buddy_sid'); if (s) return s
      const n = genId(); sessionStorage.setItem('buddy_sid', n); return n
    }
    return genId()
  })
  const bottomRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    api.get<{ data?: ChatMessage[] }>(`/chat/${sessionId}`)
      .then(d => {
        if (d.data && d.data.length > 0) setMessages(d.data)
        else setMessages([{ id: 0, session_id: sessionId, role: 'assistant',
          content: 'Xin chao! Minh la Buddy AI - nguoi ban dong hanh tam ly 24/7 cua ban. Hay chia se bat cu dieu gi ban muon nhe!',
          created_at: new Date().toISOString() }])
      })
      .catch(() => setMessages([{ id: 0, session_id: sessionId, role: 'assistant',
        content: 'Xin chao! Minh la Buddy AI. Ban dang cam thay the nao hom nay?',
        created_at: new Date().toISOString() }]))
  }, [sessionId])

  useEffect(() => { bottomRef.current?.scrollIntoView({ behavior: 'smooth' }) }, [messages])

  const handleSend = async () => {
    if (!input.trim() || loading) return
    const userMsg: ChatMessage = { id: Date.now(), session_id: sessionId, role: 'user', content: input, created_at: new Date().toISOString() }
    setMessages(p => [...p, userMsg])
    const sentInput = input 
    setInput('')
    setLoading(true)
    try {
      const data = await api.post<{ response?: string }>('/chat', { session_id: sessionId, message: sentInput })
      const reply = data.response || 'Co loi xay ra. Vui long thu lai.'
      setMessages(p => [...p, { id: Date.now() + 1, session_id: sessionId, role: 'assistant', content: reply, created_at: new Date().toISOString() }])
    } catch (err) {
      const msg = err instanceof ApiError ? err.message : 'Xin loi, co loi xay ra. Vui long thu lai!'
      setMessages(p => [...p, { id: Date.now() + 1, session_id: sessionId, role: 'assistant', content: msg, created_at: new Date().toISOString() }])
    } finally { setLoading(false) }
  }

  const quickReplies = ['Toi dang bi stress vi thi cu', 'Toi kho tap trung khi hoc', 'Toi cam thay lo lang', 'Cho toi meo hoc tap', 'Toi mat dong luc']

  return (
    <div className="max-w-4xl mx-auto px-4 py-8">
      <div className="text-center mb-8">
        <div className="text-6xl mb-3"></div>
        <h1 className="text-4xl font-bold text-gray-900 mb-2">Buddy AI</h1>
        <p className="text-gray-600">Nguoi ban dong hanh tam ly 24/7</p>
        <div className="flex items-center justify-center gap-2 mt-3">
          <div className="w-2 h-2 bg-green-500 rounded-full animate-pulse" />
          <span className="text-green-600 text-sm font-medium">Dang hoat dong</span>
        </div>
      </div>
      <div className="bg-white rounded-2xl shadow-xl border border-gray-100 overflow-hidden">
        <div className="bg-gradient-to-r from-purple-600 to-indigo-700 text-white p-4 flex items-center gap-3">
          <div className="w-10 h-10 bg-white/20 rounded-full flex items-center justify-center text-xl"></div>
          <div><div className="font-bold">Buddy AI</div><div className="text-purple-200 text-xs">Ho tro tam ly hoc sinh</div></div>
          <div className="ml-auto flex items-center gap-1">
            <div className="w-2 h-2 bg-green-400 rounded-full animate-pulse" />
            <span className="text-xs text-purple-200">Online</span>
          </div>
        </div>
        <div className="h-96 overflow-y-auto p-4 bg-gray-50 space-y-4">
          {messages.map((msg) => (
            <div key={msg.id} className={"flex " + (msg.role === 'user' ? 'justify-end' : 'justify-start')}>
              {msg.role === 'assistant' && <div className="w-8 h-8 bg-purple-100 rounded-full flex items-center justify-center text-sm mr-2 mt-1 flex-shrink-0"></div>}
              <div className={"max-w-xs md:max-w-md px-4 py-3 rounded-2xl text-sm whitespace-pre-wrap " + (msg.role === 'user' ? 'bg-blue-600 text-white rounded-br-sm' : 'bg-white text-gray-800 shadow-sm border border-gray-100 rounded-bl-sm')}>
                {msg.content}
              </div>
              {msg.role === 'user' && <div className="w-8 h-8 bg-blue-100 rounded-full flex items-center justify-center text-sm ml-2 mt-1 flex-shrink-0"></div>}
            </div>
          ))}
          {loading && (
            <div className="flex items-start gap-2">
              <div className="w-8 h-8 bg-purple-100 rounded-full flex items-center justify-center"></div>
              <div className="bg-white px-4 py-3 rounded-2xl shadow-sm border">
                <div className="flex gap-1">
                  <span className="w-2 h-2 bg-gray-400 rounded-full animate-bounce" style={{ animationDelay: '0ms' }} />
                  <span className="w-2 h-2 bg-gray-400 rounded-full animate-bounce" style={{ animationDelay: '150ms' }} />
                  <span className="w-2 h-2 bg-gray-400 rounded-full animate-bounce" style={{ animationDelay: '300ms' }} />
                </div>
              </div>
            </div>
          )}
          <div ref={bottomRef} />
        </div>
        <div className="px-4 py-2 border-t border-gray-100 bg-white">
          <p className="text-xs text-gray-400 mb-2">Goi y nhanh:</p>
          <div className="flex flex-wrap gap-2">
            {quickReplies.map(qr => (
              <button key={qr} onClick={() => setInput(qr)}
                className="text-xs bg-purple-50 text-purple-700 border border-purple-200 px-3 py-1 rounded-full hover:bg-purple-100 transition">{qr}</button>
            ))}
          </div>
        </div>
        <div className="p-4 bg-white border-t">
          <div className="flex gap-3">
            <input type="text" value={input} onChange={e => setInput(e.target.value)}
              onKeyDown={e => { if (e.key === 'Enter' && !e.shiftKey) { e.preventDefault(); handleSend() } }}
              placeholder="Nhap tin nhan... (Enter de gui)"
              className="flex-1 border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-purple-300"
              disabled={loading} />
            <button onClick={handleSend} disabled={loading || !input.trim()}
              className="bg-purple-600 text-white px-6 py-3 rounded-xl hover:bg-purple-700 disabled:opacity-50 transition font-medium text-sm">
              Gui
            </button>
          </div>
        </div>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mt-10">
        {[{ icon: '', title: 'An toan va Bao mat', desc: 'Moi cuoc tro chuyen duoc bao mat hoan toan.' },
          { icon: '', title: 'Ho tro 24/7', desc: 'Buddy AI luon san sang lang nghe bat ky luc nao.' },
          { icon: '', title: '50+ Kich ban', desc: 'Duoc huan luyen voi 50+ kich ban tam ly thuc te.' }].map(c => (
          <div key={c.title} className="bg-white rounded-xl p-5 shadow-md text-center">
            <div className="text-3xl mb-3">{c.icon}</div>
            <h3 className="font-bold text-gray-900 mb-2">{c.title}</h3>
            <p className="text-gray-500 text-sm">{c.desc}</p>
          </div>
        ))}
      </div>
    </div>
  )
}
