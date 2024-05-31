import { useState } from 'react'
import "./output.css"
import './App.css'
import { Login } from './pages/Login'

function App() {

  return (
    <>
     <div className=' lg:max-h-screen lg:mt-10'>
      <div className='lg:w-3/4 lg:mx-auto'>
        <Login/>
      </div>
     </div>
    </>
  )
}

export default App
