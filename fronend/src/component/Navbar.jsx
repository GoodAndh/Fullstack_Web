import {useAuth} from "../utils/AuthContext.jsx"

function Navbar() {
    const {isLogin}=useAuth() 
    if (isLogin === undefined) {
      // Lakukan penanganan jika isLogin masih undefined
      // Contoh: set nilai default atau tindakan lain yang diperlukan
      console.log("kontol islogin undefine ajg");
    }
    
    const handleLogOut = () => {
        localStorage.removeItem("access_token");
        localStorage.removeItem("refresh_token");
      };


  return (
   <>
    <div className="block">
        <div className="max-w-7xl mx-auto  px-4 sm:px-6 lg:px-8 bg-sky-200 rounded-2xl">
          <div className="">
            <div className="flex items-center justify-between h-12">
              <div className="hidden md:flex items-center">
                <div className="flex-shrink-0">
                  <a href="/" className="font-semibold">
                    Home
                  </a>
                </div>
              </div>
              <div className="w-full mt-3 mb-2">
                <div className=" flex mb-2 max-w-xl md:mx-auto">
                  <button className="border border-slate-800 rounded-l-md font-semibold bg-white">
                    Cari
                  </button>
                  <input
                    placeholder="cari produk.."
                    type="text"
                    className=" border border-slate-800 w-full rounded-r-lg text-center"
                  />
                </div>
              </div>
              {/*  */}

              <div className="block">
                <div className="ml-4 flex items-center space-x-4">
                  {isLogin ? (
                    <>
                      {" "}
                     
                      <a
                      href="/login"
                        type="button"
                        onClick={handleLogOut}
                        className="font-semibold focus:underline underline-offset-4 decoration-white hover:scale-125 p-2 transition"
                      >
                        LogOut
                      </a>
                      <a
                        href="/me"
                        className="hidden md:block font-semibold focus:underline underline-offset-4 decoration-white hover:scale-125 p-2 transition"
                      >
                        Profile
                      </a>
                    </>
                  ) : (
                    <>
                      {" "}
                      <a
                        href="/login"
                        className="font-semibold focus:underline underline-offset-4 decoration-white hover:scale-125 p-2 transition"
                      >
                        Masuk
                      </a>
                      <a
                        href="/register"
                        className="hidden md:block font-semibold focus:underline underline-offset-4 decoration-white hover:scale-125 p-2 transition"
                      >
                        Daftar
                      </a>
                    </>
                  )}
                </div>
              </div>
            </div>
          </div>
        </div>
        {/*  */}
      </div>
   </>

  )
}

export default Navbar