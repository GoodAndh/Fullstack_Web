import Login from "./pages/Login.jsx";
import Register from "./pages/Register.jsx";
import HomePage from "./pages/HomePage.jsx";
import ProductPage from "./pages/ProductPage.jsx";
import ProfilePage from "./pages/ProfilePage.jsx";
// import UpdateProfile from "./pages/UpdateProfile.jsx"

import { createBrowserRouter, RouterProvider } from "react-router-dom";

function App() {
  const router = createBrowserRouter([
    { path: "*", element: <div>NOT FOUND KONTOL</div> },
    { path: "/login", element: <Login /> },
    { path: "/register", element: <Register /> },
    { path: "/", element: <HomePage /> },
    { path: "/p/:pID", element: <ProductPage /> },
    { path: "/me", element: <ProfilePage /> },
    // {path:"/update/profile",element:<UpdateProfile/>}
  ]);

  return (
    <>
      <RouterProvider router={router} />
    </>
  );
}

export default App;
