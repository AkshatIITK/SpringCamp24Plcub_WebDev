import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import './App.css';
import Layout from "./components/Layout";
import Home from "./components/home";
import Blogs from "./components/blogs";
import NoPage from "./components/nopage";
import Blog from "./components/blog";
import Login from "./Pages/Login";
import Signup from "./Pages/Signup";
import { useEffect, useState } from "react";
import SubscribedBlog from "./components/SubcribeBlogs";
// import NewHome from "./components/newHome";

function App() {
  const [isLogged, setIsLogged] = useState(false);
  const [user, setUser] = useState({});
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const checkAuthentication = async () => {
      try {
        const response = await fetch('http://localhost:8080/user', {
          headers: {'Content-Type': 'application/json'},
          credentials: 'include',
        });
        if (response.ok) {
          const userRes = await response.json();
          setUser(userRes.claims); 
          setIsLogged(true);
        } else {
          setIsLogged(false);
        }
      } catch (error) {
        console.error("Error checking authentication:", error);
        setIsLogged(false);
      } finally {
        setLoading(false);
      }
    };
    checkAuthentication();
  }, []);

  if (loading) {
    // Show loading indicator while checking authentication
    return <div>Loading...</div>;
  }

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login isloginUS={setIsLogged} setUser={setUser} />} />
        <Route path="/signup" element={<Signup userUS={setUser} />} />
        {isLogged ? (
          <>
            <Route path="/" element={<Layout />}>
              <Route index element={<Home  />} />
              <Route path="/blogs" element={<Blogs user={user}/>} />
              <Route path="/blogs/:blogid" element={<Blog />} />
              {/* <Route path="/SubscribedBlogs" element={<SubscribedBlog user={user} />} /> */}
              {/* <Route path="/SubscribedBlogs/:blogid" element={<SubscribedBlog />} /> */}
              <Route path="*" element={<NoPage />} />
            </Route>
          </>
        ) : (
          <Route path="*" element={<Navigate to="/login" />} />
        )}
      </Routes>
    </BrowserRouter>
  );
}

export default App;
