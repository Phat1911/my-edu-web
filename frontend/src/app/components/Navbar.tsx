'use client'
import Link from 'next/link'
import { useState, useEffect } from 'react'
import { useRouter } from 'next/navigation'
import { isLoggedIn, getUser, logout } from '@/lib/auth'

export default function Navbar() {
  const router = useRouter()
  const [open, setOpen] = useState(false)
  const [loggedIn, setLoggedIn] = useState(false)
  const [displayName, setDisplayName] = useState('')

  useEffect(() => {
    const status = isLoggedIn()
    setLoggedIn(status)
    if (status) {
      const user = getUser()
      setDisplayName(user?.display_name || user?.username || '')
    }
  }, [])

  function handleLogout() {
    logout()
    setLoggedIn(false)
    setDisplayName('')
    router.push('/')
  }

  const links = [
    { href: '/', label: 'Trang chu', icon: '🏠' },
    { href: '/videos', label: 'Video hoc tap', icon: '🎬' },
    { href: '/audios', label: 'Am thanh song nao', icon: '🎵' },
    { href: '/chatbot', label: 'Buddy AI', icon: '🤖' },
    { href: '/qrcodes', label: 'Ma QR', icon: '📱' },
    { href: '/messages', label: 'Tin nhan', icon: '💬' },
  ]

  return (
    <nav className="bg-gradient-to-r from-blue-700 to-indigo-800 text-white shadow-lg sticky top-0 z-50">
      <div className="max-w-7xl mx-auto px-4">
        <div className="flex items-center justify-between h-16">
          <Link href="/" className="flex items-center gap-2 font-bold text-xl">
            <span className="text-2xl">📚</span>
            <span>EduHub</span>
          </Link>
          <div className="hidden md:flex items-center gap-1">
            {links.map((l) => (
              <Link key={l.href} href={l.href}
                className="flex items-center gap-1 px-3 py-2 rounded-lg text-sm font-medium hover:bg-white/20 transition-colors">
                <span>{l.icon}</span>
                <span>{l.label}</span>
              </Link>
            ))}
          </div>
          <div className="hidden md:flex items-center gap-2">
            {loggedIn ? (
              <>
                <span className="text-sm text-blue-200 px-2">
                  Xin chao, <span className="text-white font-semibold">{displayName}</span>
                </span>
                <button
                  onClick={handleLogout}
                  className="px-3 py-1.5 text-sm font-medium bg-white/20 hover:bg-white/30 rounded-lg transition-colors"
                >
                  Dang xuat
                </button>
              </>
            ) : (
              <>
                <Link href="/login"
                  className="px-3 py-1.5 text-sm font-medium hover:bg-white/20 rounded-lg transition-colors">
                  Dang nhap
                </Link>
                <Link href="/register"
                  className="px-3 py-1.5 text-sm font-medium bg-white text-blue-700 hover:bg-blue-50 rounded-lg transition-colors">
                  Dang ky
                </Link>
              </>
            )}
          </div>
          <button onClick={() => setOpen(!open)} className="md:hidden p-2 rounded-lg hover:bg-white/20">
            {open ? '✕' : '☰'}
          </button>
        </div>
        {open && (
          <div className="md:hidden pb-3 flex flex-col gap-1">
            {links.map((l) => (
              <Link key={l.href} href={l.href} onClick={() => setOpen(false)}
                className="flex items-center gap-2 px-4 py-2 rounded-lg hover:bg-white/20 text-sm font-medium">
                <span>{l.icon}</span><span>{l.label}</span>
              </Link>
            ))}
            <div className="border-t border-white/20 mt-2 pt-2 flex flex-col gap-1">
              {loggedIn ? (
                <>
                  <span className="px-4 py-1 text-sm text-blue-200">
                    Xin chao, <span className="text-white font-semibold">{displayName}</span>
                  </span>
                  <button
                    onClick={() => { handleLogout(); setOpen(false) }}
                    className="text-left px-4 py-2 rounded-lg hover:bg-white/20 text-sm font-medium"
                  >
                    Dang xuat
                  </button>
                </>
              ) : (
                <>
                  <Link href="/login" onClick={() => setOpen(false)}
                    className="px-4 py-2 rounded-lg hover:bg-white/20 text-sm font-medium">
                    Dang nhap
                  </Link>
                  <Link href="/register" onClick={() => setOpen(false)}
                    className="px-4 py-2 rounded-lg hover:bg-white/20 text-sm font-medium">
                    Dang ky
                  </Link>
                </>
              )}
            </div>
          </div>
        )}
      </div>
    </nav>
  )
}
