import React from "react";
import { useParams } from "react-router-dom";
import { getApi } from "../api/useApi.js";
import CardImage from "../component/CardImage.jsx";
import Navbar from "../component/Navbar.jsx";
import Botbar from "../component/Botbar.jsx";

function ProductPage() {
  const { pID } = useParams();
  const [error, setError] = React.useState(false);
  const [data, setData] = React.useState({
    category: "",
    deskripsi: "",
    id: null,
    name: "",
    price: null,
    quantity: null,
    url_image: "",
    userid: null,
  });

  React.useEffect(() => {
    async function callApi() {
      const { response, error } = await getApi(`apr/${pID}`);
      if (response) {
        setData(response.data);
      }
      if (error) {
        setError(true);
      }
    }
    callApi();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  if (error) {
    return <div>BARANGNYA BELUM ADA</div>;
  }

  return (
    <>
      <div className="sm:mb-20 md:mb-32 mt-1 md:mt-3">
        <Navbar />
      </div>
      {data && (
        <div className=" mb-96 flex items-center  min-h-screen  container mx-auto mt-12 bg-slate-400 rounded-xl">
          {/* START A */}
          <div className="grid grid-cols-1 mx-auto">
            {/* CONTENT A*/}
            <div className="m-2 rounded-xl shadow-lg bg-slate-300">
              <div className="p-8 flex flex-col">
                <div className="rounded-xl overflow-hidden">
                  {data.url_image && <CardImage url={data.url_image} />}
                </div>
                <h5 className="text-2xl md:text-3xl font-medium mt-3">{`${
                  data.name && data.name
                }`}</h5>
                <p className="text-md font-semibold">{`${
                  data.category && data.category
                }`}</p>
              </div>
            </div>
            {/* END CONTENT A */}
            {/*  CONTENT B */}
            <div className="grid grid-cols-2 m-2">
              {/*  */}
              <div className="m-2 rounded-xl shadow-lg bg-slate-300">
                <div className="p-8 flex flex-col">
                  <div className="rounded-xl">
                    <h5 className="text-2xl md:text-3xl font-medium mt-3">
                      Rp. {`${data.price && data.price}`}
                    </h5>

                    <p className="font-semibold  m-1">Jumlah pesanan:</p>
                    <input
                      type="number"
                      className="w-full text-center rounded-md"
                      min="0"
                      placeholder="jumlah pesanan"
                    />
                  </div>
                </div>
              </div>
              {/* SAMPING B */}
              <div className="m-2 rounded-xl shadow-lg bg-slate-300">
                <div className="p-8 flex flex-col">
                  <div className="rounded-xl">
                    <button className="rounded-md bg-sky-500 font-semibold text-white hover:scale-125 transition">
                      <p className="m-2">Beli Sekarang</p>
                    </button>
                    <br />
                    <button className="p-2 m-1 font-semibold  hover:scale-125 transition">
                      + Wishlist
                    </button>
                    <br />
                    <button className="p-2 m-1 font-semibold  hover:scale-125 transition">
                      + Keranjang
                    </button>
                  </div>
                </div>
              </div>
            </div>
            {/* END CONTENT B */}
            <div className="grid grid-cols-1">
              <div className="m-3 rounded-xl shadow-lg bg-slate-300">
                <div className="p-5 flex flex-col">
                  <div className="rounded-xl overflow-hidden">
                    <h5 className="text-2xl md:text-3xl font-medium mt-3">
                      Deskripsi Produk :
                    </h5>

                    <h5 className="text-xl font-medium mt-6 text-slate-500">
                      {`${data.deskripsi && data.deskripsi}`}
                    </h5>
                  </div>
                </div>
              </div>
            </div>
          </div>
          {/* END A */}
        </div>
      )}
      <div className="">
        <Botbar />
      </div>
    </>
  );
}

export default ProductPage;
