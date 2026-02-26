'use client'
import { useEffect, useState } from 'react'

interface Video {
  id: number; title: string; description: string
  drive_url: string; embed_url: string; thumbnail: string
  category: string; order: number; created_at: string
}

export default function VideosPage() {
  const [videos, setVideos] = useState<Video[]>([])
  const [loading, setLoading] = useState(true)
  const apiUrl = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1'

  useEffect(() => {
    fetch(apiUrl + '/videos')
      .then(r => r.json()).then(d => setVideos(d.data || []))
      .catch(console.error).finally(() => setLoading(false))
  }, [apiUrl])

  return (
    <div className="max-w-6xl mx-auto px-4 py-12">
      <div className="text-center mb-12">
        <div className="text-6xl mb-4">🎥</div>
        <h1 className="text-4xl font-bold text-gray-900 mb-4">Video Meo Hoc Tap</h1>
        <p className="text-xl text-gray-600 max-w-2xl mx-auto">Bo suu tap video ky thuat hoc tap hieu qua, duoc chung minh khoa hoc</p>
      </div>
      <div className="bg-white rounded-2xl shadow-lg border border-blue-100 overflow-hidden mb-12">
        <div className="bg-gradient-to-r from-blue-600 to-blue-700 text-white p-4 flex items-center justify-between">
          <div className="flex items-center gap-3">
            <span className="text-2xl">📁</span>
            <div>
              <h2 className="font-bold text-lg">Thu muc Video Google Drive</h2>
              <p className="text-blue-200 text-sm">Truy cap truc tiep toan bo video hoc tap</p>
            </div>
          </div>
          <a href="https://drive.google.com/drive/folders/11qtWiDzEcHheOblUSIX_wJAlptEdzyT8"
             target="_blank" rel="noopener noreferrer"
             className="bg-white text-blue-700 font-semibold px-4 py-2 rounded-lg hover:bg-blue-50 transition text-sm">
            Mo Drive
          </a>
        </div>
        <div className="relative" style={{paddingBottom:'60%'}}>
          <iframe src="https://drive.google.com/embeddedfolderview?id=11qtWiDzEcHheOblUSIX_wJAlptEdzyT8#list"
            className="absolute inset-0 w-full h-full" frameBorder="0" allowFullScreen />
        </div>
      </div>
      <h2 className="text-2xl font-bold text-gray-900 mb-6">Video Noi bat</h2>
      {loading ? (
        <div className="text-center py-12 text-gray-500">
          <div className="text-4xl mb-3 animate-pulse">🎥</div>
          <p>Dang tai video...</p>
        </div>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {videos.map((v) => (
            <div key={v.id} className="bg-white rounded-xl shadow-md hover:shadow-lg transition border border-gray-100 overflow-hidden group">
              <div className="bg-gradient-to-br from-blue-100 to-indigo-100 h-40 flex items-center justify-center text-6xl">🎥</div>
              <div className="p-5">
                <span className="bg-blue-100 text-blue-700 text-xs font-semibold px-2 py-1 rounded-full mb-3 inline-block">{v.category}</span>
                <h3 className="font-bold text-gray-900 mb-2 group-hover:text-blue-700">{v.title}</h3>
                <p className="text-gray-500 text-sm mb-4">{v.description}</p>
                <a href={v.drive_url} target="_blank" rel="noopener noreferrer"
                   className="text-blue-600 hover:text-blue-800 text-sm font-medium">Xem video</a>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
