import Navbar from './components/layout/Navbar'
import Hero from './components/sections/Hero'
import Products from './components/sections/Products'
import Footer from './components/layout/Footer'
import ChatWidget from './components/chat/ChatWidget'

function App() {
  return (
    <div className="min-h-screen bg-gray-900 font-sans">
      <Navbar />
      <main>
        <Hero />
        <Products />
      </main>
      <Footer />
      <ChatWidget />
    </div>
  )
}

export default App