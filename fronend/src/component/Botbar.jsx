
function Botbar() {
    
    return (
      <div>
        <div className="block md:hidden">
          <div className="fixed bottom-0 left-0 right-0 bg-sky-200 flex justify-around items-center py-4">
            <div className="flex flex-col items-center">
              <a
                href="/"
                className=" font-semibold focus:underline underline-offset-4 decoration-white hover:scale-125 p-2 transition"
              >
                Home
              </a>
            </div>
            <div className="flex flex-col items-center">
              <a
                href="/wishlist"
                className=" font-semibold focus:underline underline-offset-4 decoration-white hover:scale-125 p-2 transition"
              >
                Wishlist
              </a>
            </div>
            <div className="flex flex-col items-center">
              <a
                href="/cart"
                className=" font-semibold focus:underline underline-offset-4 decoration-white hover:scale-125 p-2 transition"
              >
                Keranjang
              </a>
            </div>
            <div className="flex flex-col items-center">
              <a
                href="/me"
                className=" font-semibold focus:underline underline-offset-4 decoration-white hover:scale-125 p-2 transition"
              >
                Akun
              </a>
            </div>
          </div>
        </div>
      </div>
    );
  }
  
  export default Botbar;
  