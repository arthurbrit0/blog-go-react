import { BrowserRouter as Router, Routes, Route, useLocation } from 'react-router-dom';
import Login from './pages/Login';
import Home from './pages/Home';
import ProtectedRoute from './components/ProtectedRoute';
import { AuthProvider } from './context/AuthContext';
import './index.css';
import Registrar from './pages/Registrar';
import Navbar from './components/Navbar';
import MeusPosts from './pages/MeusPosts';

function App() {

  const isAuthRoute = location.pathname === '/login' || location.pathname === '/registrar';

  return (
    <AuthProvider>
      <Router>
        <div className="pt-5 w-2/3 mx-auto">
        <Navbar />
          <div className="container mx-auto heig mt-8 p-8 bg-white shadow-lg rounded-lg border border-gray-200">
          <Routes>
            <Route path="/registrar" element={<Registrar />} />
            <Route path="/login" element={<Login />} />
            <Route
              path="/"
              element={
                <ProtectedRoute>
                  <Home />
                </ProtectedRoute>
              }
            />
            <Route
              path="/meusposts"
              element={
                <ProtectedRoute>
                  <MeusPosts />
                </ProtectedRoute>
              }
            />
          </Routes>
          </div>
        </div>
      </Router>
    </AuthProvider>
  );
}

export default App;
