```javascript
/**  
 * element是真实dom  
 * 右边是虚拟dom  
 * 这就是彻头彻尾的胶水代码  
 */  
createRoot(element).render(  
  <React.StrictMode>  
    <RouterProvider router={router} />  
  </React.StrictMode>);
```
