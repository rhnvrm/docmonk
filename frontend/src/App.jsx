import { useState } from 'react'
import docmonkLogo from './assets/docmonk.png'
import viteLogo from '/vite.svg'
import './App.css'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="App">
      <div>
        <a href="https://docmonk.app" target="_blank">
          <img src={docmonkLogo} className="logo docmonk" alt="Docmonk logo" />
        </a>
      </div>
      <h1>docmonk.app</h1>

      <p className="read-the-docs">
        coming soon
      </p>
    </div>
  )
}

export default App
