// import React from 'react'
import Navbar from "../component/Navbar.jsx"
import Botbar from "../component/Botbar.jsx"
import CardImage from "../component/CardImage.jsx"

import {useAuth} from "../utils/AuthContext.jsx"
import {useNavigate} from "react-router-dom"

function ProfilePage() {
    const nav=useNavigate()
    const {user}=useAuth()
    if (!user) {
        localStorage.removeItem("access_token");
        localStorage.removeItem("refresh_token");
        nav("/login")
        window.location.reload();
    }

  return (
   <>
    <div className="sm:mb-20 md:mb-28 mt-1 md:mt-3">
        <Navbar />
      </div>
   <div className=" mb-96 flex items-center min-h-screen  container mx-auto mt-12 bg-slate-400 rounded-xl">
        <div className="grid grid-cols-1 mx-auto">
            <div className="m-2 rounded-xl shadow-lg bg-slate-300">
                <div className="p-8 flex flex-col">
                    <div className="rounded-xl overflow-hidden">
                        <CardImage url="Screenshot 2024-05-21 010931.png"/>
                    </div>
                </div>
            </div>
            <div className="flex ">
               <div className="m-2 rounded-xl shadow-lg bg-slate-300">
                <div className="p-8 flex flex-col ">
                    {/* <div className=""></div> */}
                    <h5 className="text-2xl">hello</h5>
                </div>
               </div>
            </div>
        </div>

      

        

   </div>
   <div className="">
        <Botbar/>
      </div>
   </>
  )
}

export default ProfilePage