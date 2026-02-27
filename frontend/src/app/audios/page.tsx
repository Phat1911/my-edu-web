'use client'
import { useEffect, useState } from 'react'
import api from '@/lib/api'

interface Audio {
  id: number; title: string; description: string
  drive_url: string; embed_url: string; category: string
  duration: string; order: number; created_at: string
}

export default function AudiosPage() {
  const [audios, setAudios] = useState<Audio[]>([])
  const [loading, setLoading] = useState(true)
  const [playing, setPlaying] = useState<number | null>(null)

  useEffect(() => {
    api.get<{ data?: Audio[] }>('/audios')
      .then(d => setAudios(d.data || []))
      .catch(console.error)
      .finally(() => setLoading(false))
  }, [])

  const brainwaves = [
    { type: 'Delta', hz: '0.5-4 Hz', icon: '', desc: 'Ngu sau, hoi phuc', color: 'from-indigo-500 to-blue-600' },
    { type: 'Theta', hz: '4-8 Hz', icon: '', desc: 'Sang tao, thien dinh', color: 'from-purple-500 to-indigo-600' },
    { type: 'Alpha', hz: '8-13 Hz', icon: '', desc: 'Thu gian, tap trung', color: 'from-green-500 to-teal-600' },
    { type: 'Beta', hz: '13-30 Hz', icon: '', desc: 'Tinh tao, hoc tap', color: 'from-yellow-500 to-orange-500' },
    { type: 'Gamma', hz: '30-100 Hz', icon: '', desc: 'Hieu suat cao', color: 'from-red-500 to-pink-600' },
  ]

  return (
    <div className="max-w-6xl mx-auto px-4 py-12">
      <div className="text-center mb-12">
        <div className="text-6xl mb-4"></div>
        <h1 className="text-4xl font-bold text-gray-900 mb-4">Am thanh Song Nao</h1>
        <p className="text-xl text-gray-600 max-w-2xl mx-auto">Binaural beats giup toi uu hoa trang thai nao bo cho viec hoc tap</p>
      </div>
      <div className="mb-12">
        <h2 className="text-2xl font-bold text-gray-900 mb-6 text-center">Huong dan Song Nao</h2>
        <div className="grid grid-cols-2 md:grid-cols-5 gap-4">
          {brainwaves.map((b) => (
            <div key={b.type} className={"bg-gradient-to-br " + b.color + " text-white rounded-2xl p-4 text-center shadow-md"}>
              <div className="text-3xl mb-2">{b.icon}</div>
              <div className="font-bold text-lg">{b.type}</div>
              <div className="text-xs opacity-80 mb-1">{b.hz}</div>
              <div className="text-xs opacity-90">{b.desc}</div>
            </div>
          ))}
        </div>
      </div>
      <div className="bg-white rounded-2xl shadow-lg border border-green-100 overflow-hidden mb-12">
        <div className="bg-gradient-to-r from-green-600 to-teal-700 text-white p-4 flex items-center justify-between">
          <div className="flex items-center gap-3">
            <span className="text-2xl"></span>
            <div>
              <h2 className="font-bold text-lg">Thu muc Am thanh Google Drive</h2>
              <p className="text-green-200 text-sm">Toan bo am thanh song nao</p>
            </div>
          </div>
          <a href="https://drive.google.com/drive/folders/1tsyTAwnZyd0QwtamQsk46ZvdfY8YM0_Q"
             target="_blank" rel="noopener noreferrer"
             className="bg-white text-green-700 font-semibold px-4 py-2 rounded-lg text-sm">Mo Drive</a>
        </div>
        <div className="relative" style={{ paddingBottom: '50%' }}>
          <iframe src="https://drive.google.com/embeddedfolderview?id=1tsyTAwnZyd0QwtamQsk46ZvdfY8YM0_Q#list"
            className="absolute inset-0 w-full h-full" frameBorder="0" allowFullScreen />
        </div>
      </div>
      <h2 className="text-2xl font-bold text-gray-900 mb-6">Am thanh Noi bat</h2>
      {loading ? (
        <div className="text-center py-12 text-gray-500">
          <div className="text-4xl mb-3 animate-pulse"></div>
          <p>Dang tai am thanh...</p>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {audios.map((a) => (
            <div key={a.id} className="bg-white rounded-xl shadow-md border border-gray-100 overflow-hidden">
              <div className="bg-gradient-to-br from-green-100 to-teal-100 p-6 flex flex-col items-center">
                <div className={"text-5xl mb-3 " + (playing === a.id ? 'animate-bounce' : '')}></div>
              </div>
              <div className="p-5">
                <div className="flex items-center justify-between mb-2">
                  <span className="bg-green-100 text-green-700 text-xs font-semibold px-2 py-1 rounded-full">{a.category}</span>
                  {a.duration && <span className="text-gray-400 text-xs">{a.duration}</span>}
                </div>
                <h3 className="font-bold text-gray-900 mb-2">{a.title}</h3>
                <p className="text-gray-500 text-sm mb-4">{a.description}</p>
                <div className="flex gap-2">
                  <button onClick={() => setPlaying(playing === a.id ? null : a.id)}
                    className={"flex-1 py-2 rounded-lg text-sm font-medium transition " + (playing === a.id ? 'bg-red-100 text-red-600' : 'bg-green-100 text-green-700')}>
                    {playing === a.id ? 'Tam dung' : 'Nghe thu'}
                  </button>
                  <a href={a.drive_url} target="_blank" rel="noopener noreferrer"
                     className="bg-gray-100 text-gray-600 px-3 py-2 rounded-lg text-sm hover:bg-gray-200"></a>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
