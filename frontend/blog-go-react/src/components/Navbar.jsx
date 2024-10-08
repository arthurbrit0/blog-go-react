import { Link } from 'react-router-dom';

function Navbar() {
  return (
    <nav className="bg-white shadow-md rounded-lg w-1/2 mx-auto">
      <div className="container mx-auto px-4 py-2 flex justify-between items-center">
        <Link to="/" className="text-2xl font-bold">Blog Golang</Link>
        <div>
          <Link to="/login" className="mr-4">Login</Link>
          <Link to="/registrar">Registrar</Link>
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
