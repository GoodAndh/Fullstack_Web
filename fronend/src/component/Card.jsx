import React from "react";
import { getApi } from "../api/useApi.js";
import CardImage from "./CardImage.jsx";

function Card() {
  const [data, setData] = React.useState([]);
  

  React.useEffect(() => {
    async function callApi() {
      const { response } = await getApi("apr");
      if (response) {
        setData(response.data);
      }
    }
    callApi();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (
    <>
      <div className="flex items-center justify-center min-h-screen container mx-auto mt-12 mb-96">
        {/* GRID */}
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
          {/* CARD */}
          {data &&
            data.map(
              ({
                id,
                category,
                deskripsi,
                name,
                price,
                quantity,
                url_image,
              }) => {
                return (
                  <div key={id} className="rounded-xl shadow-lg bg-white">
                    {/* CONTENT */}
                    <div className="p-5 flex flex-col">
                      {/* IMG */}
                      <div className="rounded-xl overflow-hidden">
                        <CardImage url={url_image.length > 8 && url_image} />
                      </div>
                      <h5 className="text-2xl md:text-3xl font-medium mt-3">
                        {`${name}`}
                      </h5>
                      <p className="text-slate-800 text-lg m-3 ">
                        {`Category: ${category}`}
                      </p>
                      <p className="text-slate-800 text-lg m-3 ">
                        {`Deskripsi: ${deskripsi}`}
                      </p>
                      <p className="text-slate-800 text-lg m-3">
                        {`Stock: ${quantity}`}
                      </p>
                      <p className="text-slate-800 text-lg m-3">
                        {`Harga: ${price}`}
                      </p>
                      <a
                        href={`/p/${id}`}
                        className="text-center bg-blue-400 py-2 rounded-lg font-semibold mt-4 hover:bg-blue-300 focus:scale-95 transition-all duration-200 ease-out"
                      >
                        Selengkapnya
                      </a>
                    </div>
                    {/* END CONTENT */}
                  </div>
                );
              }
            )}
          {/* END CARD */}
        </div>
        {/* END GRID */}
      </div>
    </>
  );
}

export default Card;
