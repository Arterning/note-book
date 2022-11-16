import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);


//index.js是webpack的入口文件
//React.StrictMode用于检查组件代码不太合理的地方

//用于记录页面的性能的
reportWebVitals();
