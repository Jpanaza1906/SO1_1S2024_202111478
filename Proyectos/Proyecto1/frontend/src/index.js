import React from 'react';
import ReactDOM from 'react-dom/client';
import reportWebVitals from './reportWebVitals';
import {BrowserRouter, Routes, Route} from 'react-router-dom';
import Monitor from './pages/monitor';
import Proceso from './pages/proceso';
import Header from './components/header';
import './css/index.css'

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <div>
    <BrowserRouter>
      <Header />
      <Routes>
        <Route index element={<Monitor />} />
        <Route path='/monitor' element={<Monitor />} />
        <Route path='/proceso' element={<Proceso />} />
      </Routes>
    </BrowserRouter>
  </div>
);

reportWebVitals();
