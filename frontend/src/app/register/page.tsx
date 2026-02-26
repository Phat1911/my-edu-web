'use client'
import { useState } from 'react'
import { useRouter } from 'next/navigation'
import Link from 'next/link'
import { saveToken, saveUser } from '@/lib/auth'

const API_BASE = 'http://localhost:8080'

export default function RegisterPage() {
  const router = useRouter()
  const [form, setForm] = useState({ username: '', email: '', display_name: '', password: '' })
  const [error, setError] = useState('')
  const [loading, setLoading] = useState(false)

  function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
    setForm(prev => ({ ...prev, [e.target.name]: e.target.value }))
  }

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    setError('')
    setLoading(true)
    try {
      const res = await fetch(`${API_BASE}/api/v1/auth/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(form),
      })
      const data = await res.json()
      if (!res.ok) {
        setError(data.message || data.error || 'Dang ky that bai')
        return
      }
      saveToken(data.token)
      saveUser(data.user)
      router.push('/')
    } catch {
      setError('Loi ket noi. Vui long thu lai.')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center py-12 px-4">
      <div className="w-full max-w-md">
        <div className="bg-gradient-to-br from-blue-700 to-indigo-800 rounded-2xl p-8 text-white text-center mb-6 shadow-lg">
          <div className="text-4xl mb-2"></div>
          <h1 className="text-2xl font-bold">Tao tai khoan EduHub</h1>
          <p className="text-blue-200 text-sm mt-1">Bat dau hanh trinh hoc tap cua ban</p>
        </div>
        <div className="bg-white rounded-2xl shadow-md p-8">
          {error && (
            <div className="mb-4 p-3 bg-red-50 border border-red-200 text-red-700 rounded-lg text-sm">
              {error}
            </div>
          )}
          <form onSubmit={handleSubmit} className="space-y-5">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Ten dang nhap</label>
              <input
                type="text"
                name="username"
                value={form.username}
                onChange={handleChange}
                required
                className="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 text-gray-900"
                placeholder="Nhap ten dang nhap"
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Email</label>
              <input
                type="email"
                name="email"
                value={form.email}
                onChange={handleChange}
                required
                className="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 text-gray-900"
                placeholder="Nhap dia chi email"
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Ten hien thi</label>
              <input
                type="text"
                name="display_name"
                value={form.display_name}
                onChange={handleChange}
                required
                className="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 text-gray-900"
                placeholder="Nhap ten hien thi"
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Mat khau</label>
              <input
                type="password"
                name="password"
                value={form.password}
                onChange={handleChange}
                required
                className="w-full px-4 py-2.5 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 text-gray-900"
                placeholder="Nhap mat khau"
              />
            </div>
            <button
              type="submit"
              disabled={loading}
              className="w-full py-2.5 bg-gradient-to-r from-blue-700 to-indigo-800 text-white font-semibold rounded-lg hover:opacity-90 transition disabled:opacity-60"
            >
              {loading ? 'Dang xu ly...' : 'Dang ky'}
            </button>
          </form>
          <p className="text-center text-sm text-gray-500 mt-6">
            Da co tai khoan?{' '}
            <Link href="/login" className="text-blue-600 font-medium hover:underline">
              Dang nhap
            </Link>
          </p>
        </div>
      </div>
    </div>
  )
}
