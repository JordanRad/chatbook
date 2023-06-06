import React from 'react';
import { BrowserRouter as Router, Route, Routes, Link } from 'react-router-dom';
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import Login from './screens/Login';
function App() {
  return (
    <Router>

      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/contact" element={<Contact />} />
      </Routes>
    </Router>
  );
}

function Home() {
  return <h2>Home</h2>;
}


function Contact() {
  return <h2>Contact</h2>;
}

export default App;
