import { Link } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';
import { useNavigate } from 'react-router-dom';

function Navbar() {

  const navigate = useNavigate();

  const { auth, checkAuth } = useAuth(); // usando o hook que criamos para checkar se o usuario esta autenticado

  const handleLogout = async () => { // função para fazer logout
    try {
      const response = await fetch('http://localhost:3000/api/logout', { // fazemos um post na rota de logout
        method: 'POST',
        headers: {
          'Content-Type': 'application/json', // a rota de logout pega o cookie que tem o jwt enviado pelo usuario e o deleta
        },
        credentials: 'include',
      });

      if (response.ok) {
        console.log('Logout efetuado com sucesso');
        checkAuth(); // checamos se o usuario tem acesso ao endpoint /api/me. se não tiver, auth sera null
        navigate('/');  
        
      }
    } catch (error) {
      console.error(error);
    }
  }

  return (
    <nav className="bg-gradient-to-r from-blue-600 to-blue-800 shadow-lg rounded-lg w-9/10 mx-auto py-2">
    <div className="container mx-auto px-4 py-2 flex justify-between items-center text-white">
      <Link to="/" className="text-2xl font-bold hover:scale-105 transition-all">Blog Golang</Link>
      <div className="mx-auto justify-between space-x-4 md:space-x-8 lg:space-x-12">
        <Link to="/meusposts" className="hover:scale-110 transition-all items-center">Meus Posts</Link>
        <Link to="/criarpost" className="hover:scale-110 transition-all items-center">Criar post</Link>
      </div>
      <div className="flex">
        {auth ? ( // se o usuário estiver autenticado, mostrar o botão de logout
          <button
            onClick={handleLogout}
            className="bg-red-600 p-2 rounded-md text-white hover:scale-110 transition-all">
            Logout
          </button>
        ) : ( // se o usuário não estiver autenticado, mostrar os botões de login e registrar
          <>
            <div className="hover:scale-110 transition-all items-center">
              <Link to="/login" className="mr-4">Login</Link>
            </div>
            <div className="hover:scale-110 transition-all items-center">
              <Link to="/registrar" className="bg-white p-2 rounded-md text-black">Registrar</Link>
            </div>
          </>
        )}
      </div>
    </div>
  </nav>
  );
}

export default Navbar;
