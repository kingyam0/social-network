import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'
import './components/loggedIn/loggedIn.css'
import './components/profile/profile.css'
import './components/posts/posts.css'

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
