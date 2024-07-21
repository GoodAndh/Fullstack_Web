/* eslint-disable react-refresh/only-export-components */
/* eslint-disable react/prop-types */
import React from "react";
import  { getApi } from "../api/useApi.js";
import { accessToken } from "./store.js";

const AuthContext = React.createContext();

export const AuthProvider = ({ children }) => {
  const [isLogin, setIsLogin] = React.useState(false);
  const [user,setUser]=React.useState({})
  const myToken = accessToken;


 

  const setLogin = (bool) => {
    setIsLogin(bool);
  };

  React.useEffect(() => {

    async function callApi() {
      const { response, error } = await getApi("me");

      if (response) {
        setUser(response.data)
        setLogin(true);
      }
      if (error) {
        setLogin(false);
      }
    }

    callApi();
  }, [myToken]);


 
  

  return (
    <AuthContext.Provider value={{ setLogin, isLogin,user }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  return React.useContext(AuthContext);
};

export default AuthProvider;
