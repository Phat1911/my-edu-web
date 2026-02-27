'use client'
import { useEffect, useState } from 'react'
import api, { ApiError } from '@/lib/api'

interface QRCode {
  id: number; label: string; target_url: string
  type: string; qr_data: string; created_at: string
}

export default function QRCodesPage() {
  const [qrcodes, setQRCodes] = useState<QRCode[]>([])
  const [loading, setLoading] = useState(true)
  const [generating, setGenerating] = useState(false)
  const [form, setForm] = useState({ label: '', target_url: '', type: 'general' })
  const [showForm, setShowForm] = useState(false)

  useEffect(() => {
    api.get<{ data?: QRCode[] }>('/qrcodes')
      .then(d => setQRCodes(d.data || []))
      .catch(console.error)
      .finally(() => setLoading(false))
  }, [])

  const doGenerate = async (label: string, url: string, type: string) => {
    setGenerating(true)
    try {
      const data = await api.post<{ data?: QRCode }>('/qrcodes/generate', { label, target_url: url, type })
      if (data.data) setQRCodes(p => [data.data!, ...p])
    } catch (err) {
      console.error(err instanceof ApiError ? err.message : 'QR error')
    } finally {
      setGenerating(false)
    }
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!form.label || !form.target_url) return
    await doGenerate(form.label, form.target_url, form.type)
    setForm({ label: '', target_url: '', type: 'general' })
    setShowForm(false)
  }

  const presets = [
    { label: 'Video Hoc Tap', url: 'https://drive.google.com/drive/folders/11qtWiDzEcHheOblUSIX_wJAlptEdzyT8', type: 'video', icon: '', color: 'bg-blue-50 border-blue-200' },
    { label: 'Am thanh Song Nao', url: 'https://drive.google.com/drive/folders/1tsyTAwnZyd0QwtamQsk46ZvdfY8YM0_Q', type: 'audio', icon: '', color: 'bg-green-50 border-green-200' },
    { label: 'Buddy AI Chatbot', url: typeof window !== 'undefined' ? window.location.origin + '/chatbot' : '/chatbot', type: 'chatbot', icon: '', color: 'bg-purple-50 border-purple-200' },
  ]

  const typeColors: Record<string, string> = {
    video: 'bg-blue-100 text-blue-700', audio: 'bg-green-100 text-green-700',
    chatbot: 'bg-purple-100 text-purple-700', general: 'bg-gray-100 text-gray-700'
  }

  return (
    <div className="max-w-6xl mx-auto px-4 py-12">
      <div className="text-center mb-12">
        <div className="text-6xl mb-4"></div>
        <h1 className="text-4xl font-bold text-gray-900 mb-4">He thong Ma QR</h1>
        <p className="text-xl text-gray-600 max-w-2xl mx-auto">Tao ma QR thong minh de in vao So tay hoc sinh</p>
      </div>
      <div className="bg-white rounded-2xl shadow-md border border-gray-100 p-8 mb-10">
        <h2 className="text-xl font-bold text-gray-900 mb-6">Tao nhanh QR cho So tay</h2>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          {presets.map(p => (
            <div key={p.label} className={"rounded-xl border-2 p-5 " + p.color}>
              <div className="text-3xl mb-3">{p.icon}</div>
              <h3 className="font-bold text-gray-900 mb-1">{p.label}</h3>
              <p className="text-gray-500 text-xs mb-4 truncate">{p.url}</p>
              <button onClick={() => doGenerate(p.label, p.url, p.type)} disabled={generating}
                className="w-full bg-gray-800 text-white py-2 rounded-lg text-sm font-medium hover:bg-gray-700 transition disabled:opacity-50">
                {generating ? 'Dang tao...' : 'Tao ma QR'}
              </button>
            </div>
          ))}
        </div>
      </div>
      <div className="bg-white rounded-2xl shadow-md border border-gray-100 p-8 mb-10">
        <div className="flex items-center justify-between mb-6">
          <h2 className="text-xl font-bold text-gray-900">Tao ma QR tuy chinh</h2>
          <button onClick={() => setShowForm(!showForm)}
            className="bg-orange-500 text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-orange-600 transition">
            {showForm ? 'An form' : '+ Tao moi'}
          </button>
        </div>
        {showForm && (
          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Ten nhan *</label>
              <input type="text" value={form.label} onChange={e => setForm({ ...form, label: e.target.value })}
                placeholder="Vi du: Video Toan lop 10"
                className="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-orange-300" required />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">URL dich *</label>
              <input type="url" value={form.target_url} onChange={e => setForm({ ...form, target_url: e.target.value })}
                placeholder="https://..."
                className="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-orange-300" required />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">Loai</label>
              <select value={form.type} onChange={e => setForm({ ...form, type: e.target.value })}
                className="w-full border border-gray-200 rounded-xl px-4 py-3 text-sm focus:outline-none focus:ring-2 focus:ring-orange-300">
                <option value="general">Chung</option>
                <option value="video">Video</option>
                <option value="audio">Am thanh</option>
                <option value="chatbot">Chatbot</option>
              </select>
            </div>
            <button type="submit" disabled={generating}
              className="w-full bg-orange-500 text-white py-3 rounded-xl font-semibold hover:bg-orange-600 transition disabled:opacity-50">
              {generating ? 'Dang tao...' : 'Tao ma QR'}
            </button>
          </form>
        )}
      </div>
      <h2 className="text-2xl font-bold text-gray-900 mb-6">Danh sach Ma QR</h2>
      {loading ? (
        <div className="text-center py-12 text-gray-500">
          <div className="text-4xl mb-3 animate-pulse"></div><p>Dang tai ma QR...</p>
        </div>
      ) : qrcodes.length === 0 ? (
        <div className="text-center py-16 bg-white rounded-2xl border-2 border-dashed border-gray-200">
          <div className="text-5xl mb-4"></div>
          <h3 className="text-lg font-semibold text-gray-700 mb-2">Chua co ma QR nao</h3>
          <p className="text-gray-400 text-sm">Tao nhanh bang cac nut ben tren</p>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {qrcodes.map(qr => (
            <div key={qr.id} className="bg-white rounded-2xl shadow-md border border-gray-100 overflow-hidden hover:shadow-lg transition">
              <div className="bg-white p-6 flex items-center justify-center border-b">
                {qr.qr_data ? (
                  <img src={qr.qr_data} alt={qr.label} className="w-40 h-40 object-contain" />
                ) : (
                  <div className="w-40 h-40 bg-gray-100 rounded-xl flex items-center justify-center text-4xl"></div>
                )}
              </div>
              <div className="p-5">
                <div className="flex items-center justify-between mb-2">
                  <span className={"text-xs font-semibold px-2 py-1 rounded-full " + (typeColors[qr.type] || typeColors.general)}>{qr.type}</span>
                  <span className="text-gray-400 text-xs">{new Date(qr.created_at).toLocaleDateString('vi-VN')}</span>
                </div>
                <h3 className="font-bold text-gray-900 mb-2">{qr.label}</h3>
                <p className="text-gray-400 text-xs mb-4 truncate">{qr.target_url}</p>
                <div className="flex gap-2">
                  <a href={qr.target_url} target="_blank" rel="noopener noreferrer"
                     className="flex-1 text-center bg-gray-100 text-gray-700 py-2 rounded-lg text-xs font-medium hover:bg-gray-200 transition">Mo link</a>
                  {qr.qr_data && (
                    <a href={qr.qr_data} download={"qr-" + qr.label + ".png"}
                       className="bg-orange-100 text-orange-700 px-3 py-2 rounded-lg text-xs hover:bg-orange-200 transition">Tai</a>
                  )}
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
      <div className="bg-gradient-to-br from-orange-50 to-yellow-50 rounded-2xl p-8 mt-12">
        <h3 className="text-xl font-bold text-gray-900 mb-4">Huong dan in vao So tay</h3>
        <ol className="space-y-3">
          {['Tao ma QR cho Video, Audio va Chatbot bang nut Tao nhanh', 'Tai xuong tung ma QR bang nut Tai', 'In ra giay, cat va dan vao So tay hoc sinh', 'Dung camera dien thoai quet ma de truy cap ngay'].map((s, i) => (
            <li key={i} className="flex items-start gap-3">
              <span className="w-6 h-6 bg-orange-500 text-white rounded-full flex items-center justify-center text-xs font-bold flex-shrink-0">{i + 1}</span>
              <span className="text-gray-700 text-sm">{s}</span>
            </li>
          ))}
        </ol>
      </div>
    </div>
  )
}
