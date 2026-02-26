'use client'
import Link from 'next/link'
import { useState } from 'react'

export default function Navbar() {
  const [open, setOpen] = useState(false)

  const links = [
    { href: '/', label: 'Trang chu', icon: '🏠' },
    { href: '/videos', label: 'Video hoc tap', icon: '🎥' },
    { href: '/audios', label: 'Am thanh song nao', icon: '🎵' },
    { href: '/chatbot', label: 'Buddy AI', icon: '🤖' },
    { href: '/qrcodes', label: 'Ma QR', icon: '📱' },
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
          </div>
        )}
      </div>
    </nav>
  )
}
