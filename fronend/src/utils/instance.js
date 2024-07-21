import axios from "axios";
import { accessToken,refreshToken } from "./store.js";



// refreshToken
const instance = axios.create({
  baseURL: "http://localhost:3000/api/v1/",
});

export const refreshTokenFunc=async()=>{
  try {
    const re=await instance.post("refresh-token",null,{
      headers:{
        refresh_token:refreshToken
      }
    })
    return re.data
  } catch (error) {
    return Promise.reject(error)
  }
}

instance.interceptors.request.use(
  function (config) {
    config.headers[`Authorization`] = accessToken;
    return config;
  },
  function (error) {
    return Promise.reject(error);
  }
);


instance.interceptors.response.use(
    function (res) {
        return res;
    },
    async function (error) {
        try {
          if (error.response&&error.response.status===401) {
              await refreshTokenFunc()
            }
           return Promise.reject(error);
          } catch  {
            return Promise.reject(error);

        }
    }

)

export default instance