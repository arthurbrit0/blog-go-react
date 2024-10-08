function Home() {
    return (
      <div className="container mx-auto px-4 py-6">
        <h1 className="text-3xl font-bold mb-4">Último Posts</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div className="bg-white p-4 shadow-md rounded">
            <h2 className="text-xl font-semibold">Título do Post</h2>
            <p className="text-gray-700">Descrição do post</p>
            <a href="#" className="text-blue-500 hover:underline">Leia mais</a>
          </div>
        </div>
      </div>
    );
  }
  
  export default Home;
  