import { Link } from 'react-router-dom';

function Navbar() {
  return (
    <nav className="bg-gradient-to-r from-blue-600 to-blue-800 shadow-lg rounded-lg w-2/3 mx-auto py-2">
      <div className="container mx-auto px-4 py-2 flex justify-between items-center text-white">
        <Link to="/" className="text-2xl font-bold hover:scale-105 transition-all">Blog Golang</Link>
        <div className="flex">
          <div className="hover:scale-110 transition-all items-center">
            <Link to="/login" className="mr-4">Login</Link>
          </div>
          <div className="hover:scale-110 transition-all items-center">
            <Link to="/registrar" className="bg-white p-2 rounded-md text-black">Registrar</Link>
          </div>
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
