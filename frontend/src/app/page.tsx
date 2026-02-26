import Link from 'next/link'

export default function Home() {
  const features = [
    {
      icon: '🤖',
      title: 'Buddy AI - Chatbot Tam ly',
      description: 'Nguoi ban dong hanh 24/7, lang nghe va ho tro tam ly hoc sinh voi 50+ kich ban tu van chuyen sau. Tao khong gian an toan de chia se.',
      href: '/chatbot',
      color: 'from-purple-500 to-pink-500',
      badge: 'AI Powered',
    },
    {
      icon: '🎥',
      title: 'Video Meo Hoc Tap',
      description: 'Bo suu tap video huong dan cac ky thuat hoc tap hieu qua, ky thuat Pomodoro, tu duy tich cuc va nhieu phuong phap duoc chung minh khoa hoc.',
      href: '/videos',
      color: 'from-blue-500 to-cyan-500',
      badge: 'Kho video',
    },
    {
      icon: '🎵',
      title: 'Am thanh Song Nao',
      description: 'Bo am thanh song nao Alpha, Theta, Beta giup tang tap trung, sang tao va thu gian. Ho tro qua trinh hoc tap va ghi nho hieu qua.',
      href: '/audios',
      color: 'from-green-500 to-teal-500',
      badge: 'Brainwave',
    },
    {
      icon: '📱',
      title: 'He thong Ma QR',
      description: 'Tao ma QR thong minh cho tung tai nguyen hoc tap. In truc tiep vao So tay hoc sinh de truy cap nhanh moi luc moi noi.',
      href: '/qrcodes',
      color: 'from-orange-500 to-yellow-500',
      badge: 'Smart QR',
    },
  ]

  const steps = [
    {
      num: '01',
      title: 'Thiet lap Buddy AI',
      desc: 'Chatbot tam ly duoc nap 50 kich ban tu van, dong vai nguoi ban lang nghe 24/7',
      icon: '🤖',
    },
    {
      num: '02',
      title: 'Landing Page Tich hop',
      desc: 'Tat ca tai nguyen: am thanh song nao, video hoc tap va cua so chat AI trong mot giao dien',
      icon: '🖥️',
    },
    {
      num: '03',
      title: 'He thong Ma QR',
      desc: 'Ma QR cho tung muc duoc in vao So tay de hoc sinh truy cap nhanh',
      icon: '📱',
    },
  ]

  const stats = [
    { label: 'Video hoc tap', value: '10+', icon: '🎥' },
    { label: 'Am thanh song nao', value: '8+', icon: '🎵' },
    { label: 'Kich ban tam ly', value: '50+', icon: '💬' },
    { label: 'Ma QR thong minh', value: '∞', icon: '📱' },
  ]

  return (
    <div className="overflow-hidden">
      {/* Hero */}
      <section className="bg-gradient-to-br from-blue-700 via-indigo-800 to-purple-900 text-white py-24 px-4 relative">
        <div className="absolute inset-0 opacity-10">
          <div className="absolute top-10 left-10 w-72 h-72 bg-white rounded-full blur-3xl" />
          <div className="absolute bottom-10 right-10 w-96 h-96 bg-purple-300 rounded-full blur-3xl" />
        </div>
        <div className="max-w-5xl mx-auto text-center relative z-10">
          <div className="inline-flex items-center gap-2 bg-white/20 backdrop-blur px-4 py-2 rounded-full text-sm font-medium mb-6">
            <span>✨</span>
            <span>Nen tang Giao duc Thong minh</span>
          </div>
          <h1 className="text-5xl md:text-6xl font-bold mb-6 leading-tight">
            Hoc tap Hieu qua
            <br />
            <span className="text-yellow-300">Phat trien Toan dien</span>
          </h1>
          <p className="text-xl md:text-2xl text-blue-100 mb-10 max-w-3xl mx-auto leading-relaxed">
            Tich hop AI tam ly, video hoc tap, am thanh song nao va ma QR thong minh
            de ho tro hoc sinh phat trien toan dien ca tri tue lan cam xuc.
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Link
              href="/chatbot"
              className="bg-yellow-400 text-gray-900 font-bold px-8 py-4 rounded-xl hover:bg-yellow-300 transition-all shadow-lg text-lg"
            >
              🤖 Tro chuyen voi Buddy AI
            </Link>
            <Link
              href="/videos"
              className="bg-white/20 backdrop-blur border border-white/30 font-semibold px-8 py-4 rounded-xl hover:bg-white/30 transition-all text-lg"
            >
              🎥 Xem Video hoc tap
            </Link>
          </div>
        </div>
      </section>

      {/* Stats */}
      <section className="bg-white shadow-sm py-8 px-4">
        <div className="max-w-5xl mx-auto grid grid-cols-2 md:grid-cols-4 gap-6">
          {stats.map((s) => (
            <div key={s.label} className="text-center">
              <div className="text-3xl mb-1">{s.icon}</div>
              <div className="text-3xl font-bold text-blue-700">{s.value}</div>
              <div className="text-sm text-gray-500">{s.label}</div>
            </div>
          ))}
        </div>
      </section>

      {/* Features */}
      <section className="py-20 px-4 bg-gray-50">
        <div className="max-w-6xl mx-auto">
          <div className="text-center mb-14">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">Tinh nang noi bat</h2>
            <p className="text-xl text-gray-600 max-w-2xl mx-auto">
              Bon tru cot ho tro hoc sinh phat trien toan dien
            </p>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            {features.map((f) => (
              <Link
                key={f.href}
                href={f.href}
                className="group bg-white rounded-2xl p-8 shadow-md hover:shadow-xl transition-all border border-gray-100 hover:-translate-y-1"
              >
                <div className={"inline-flex items-center gap-2 bg-gradient-to-r " + f.color + " text-white px-3 py-1 rounded-full text-xs font-semibold mb-4"}>
                  {f.badge}
                </div>
                <div className="text-5xl mb-4">{f.icon}</div>
                <h3 className="text-xl font-bold text-gray-900 mb-3 group-hover:text-blue-700 transition-colors">
                  {f.title}
                </h3>
                <p className="text-gray-600 leading-relaxed">{f.description}</p>
                <div className="mt-4 flex items-center gap-1 text-blue-600 font-medium text-sm">
                  <span>Kham pha ngay</span>
                  <span className="group-hover:translate-x-1 transition-transform">→</span>
                </div>
              </Link>
            ))}
          </div>
        </div>
      </section>

      {/* Steps */}
      <section className="py-20 px-4 bg-white">
        <div className="max-w-5xl mx-auto">
          <div className="text-center mb-14">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">Quy trinh 3 Buoc</h2>
            <p className="text-xl text-gray-600">Don gian, hieu qua va de tiep can</p>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            {steps.map((s, i) => (
              <div key={s.num} className="relative text-center">
                {i < steps.length - 1 && (
                  <div className="hidden md:block absolute top-8 left-2/3 w-full h-0.5 bg-blue-200 z-0" />
                )}
                <div className="relative z-10 inline-flex items-center justify-center w-16 h-16 bg-blue-100 text-blue-700 rounded-2xl text-3xl mb-4 shadow-md">
                  {s.icon}
                </div>
                <div className="text-xs font-bold text-blue-500 mb-2">BUOC {s.num}</div>
                <h3 className="text-lg font-bold text-gray-900 mb-3">{s.title}</h3>
                <p className="text-gray-600 text-sm leading-relaxed">{s.desc}</p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Google Drive Resources */}
      <section className="py-20 px-4 bg-gradient-to-br from-indigo-50 to-purple-50">
        <div className="max-w-5xl mx-auto">
          <div className="text-center mb-10">
            <h2 className="text-4xl font-bold text-gray-900 mb-4">Tai nguyen Google Drive</h2>
            <p className="text-gray-600 text-xl">Truy cap truc tiep kho video va am thanh</p>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div className="bg-white rounded-2xl p-8 shadow-md border border-blue-100">
              <div className="text-4xl mb-4">🎥</div>
              <h3 className="text-xl font-bold text-gray-900 mb-3">Thu muc Video Hoc Tap</h3>
              <p className="text-gray-600 mb-6 text-sm leading-relaxed">
                Bo suu tap day du video meo hoc tap, phuong phap ghi nho, tu duy sang tao va phat trien ban than.
              </p>
              <a
                href="https://drive.google.com/drive/folders/11qtWiDzEcHheOblUSIX_wJAlptEdzyT8"
                target="_blank"
                rel="noopener noreferrer"
                className="inline-flex items-center gap-2 bg-blue-600 text-white px-6 py-3 rounded-xl hover:bg-blue-700 transition-colors font-medium"
              >
                <span>📁</span>
                <span>Mo thu muc Drive</span>
              </a>
            </div>
            <div className="bg-white rounded-2xl p-8 shadow-md border border-green-100">
              <div className="text-4xl mb-4">🎵</div>
              <h3 className="text-xl font-bold text-gray-900 mb-3">Thu muc Am thanh Song Nao</h3>
              <p className="text-gray-600 mb-6 text-sm leading-relaxed">
                Toan bo bo suu tap am thanh song nao Alpha, Theta, Beta va Gamma giup toi uu hoa hieu suat hoc tap.
              </p>
              <a
                href="https://drive.google.com/drive/folders/1tsyTAwnZyd0QwtamQsk46ZvdfY8YM0_Q"
                target="_blank"
                rel="noopener noreferrer"
                className="inline-flex items-center gap-2 bg-green-600 text-white px-6 py-3 rounded-xl hover:bg-green-700 transition-colors font-medium"
              >
                <span>📁</span>
                <span>Mo thu muc Drive</span>
              </a>
            </div>
          </div>
        </div>
      </section>

      {/* CTA */}
      <section className="py-20 px-4 bg-gradient-to-r from-blue-700 to-indigo-800 text-white">
        <div className="max-w-3xl mx-auto text-center">
          <h2 className="text-4xl font-bold mb-6">Bat dau hanh trinh hoc tap ngay hom nay!</h2>
          <p className="text-xl text-blue-100 mb-8">
            Tro chuyen voi Buddy AI, xem video hoc tap hoac nghe am thanh song nao de toi uu hieu suat.
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center">
            <Link
              href="/chatbot"
              className="bg-yellow-400 text-gray-900 font-bold px-8 py-4 rounded-xl hover:bg-yellow-300 transition-all shadow-lg text-lg"
            >
              🤖 Chat voi Buddy AI ngay
            </Link>
            <Link
              href="/qrcodes"
              className="border-2 border-white font-semibold px-8 py-4 rounded-xl hover:bg-white/10 transition-all text-lg"
            >
              📱 Tao ma QR cho so tay
            </Link>
          </div>
        </div>
      </section>
    </div>
  )
}
