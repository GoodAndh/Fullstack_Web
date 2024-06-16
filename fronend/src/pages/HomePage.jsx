import Navbar from "../component/Navbar.jsx";
import Botbar from "../component/Botbar.jsx";
import Card from "../component/Card.jsx";

function HomePage() {
  return (
    <>
      <div className="sm:mb-20 md:mb-32 mt-1 md:mt-3">
        <Navbar />
      </div>
      <div className="mx-2">
        <Card />
      </div>
      <div className=" mt-36">
        <Botbar />
      </div>
    </>
  );
}

export default HomePage;
