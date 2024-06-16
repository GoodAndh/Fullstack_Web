import Form from "../component/Form.jsx"; // inputID, title, Type, OnInputChange, validate
import Button from "../component/Button.jsx";
import Navbar from "../component/Navbar.jsx";
import Botbar from "../component/Botbar.jsx";
import React from "react";
import { postApi } from "../api/useApi.js";
import { saveToken } from "../utils/store.js";
import {useNavigate} from "react-router-dom"

function Login() {
  const nav =useNavigate()
  const [formValue, setFormValue] = React.useState({
    username: "",
    password: "",
  });
  const [error, setError] = React.useState("");

  const inputChange = (id, value) => {
    setFormValue((p) => ({
      ...p,
      [id]: value,
    }));
  };
  const handleSubmit = async (e) => {
    const { error, response } = await postApi("login",formValue);
    if (error) {
      setError(error.message);
    }
    if (response) {
      // console.log("sukses login:",response)
      setError("");
      Object.entries(response.data).map(([key, value]) => {
        saveToken(key, value);
      });
      nav("/")
      window.location.reload();
    }
    e.preventDefault;
  };
  return (
    <>
      <div className="sm:mb-20 md:mb-32 mt-1 md:mt-3">
        <Navbar />
      </div>
      <div className="max-w-lg m-3 mx-auto p-5 border-2 border-slate-600 rounded-xl shadow-md  ">
        <Form
          inputID={"username"}
          title={"Username"}
          Type={"text"}
          OnInputChange={inputChange}
          validate={error}
        />
        <Form
          inputID={"password"}
          title={"Password"}
          Type={"password"}
          OnInputChange={inputChange}
          validate={error}
        />
        <div className="m-1">
        <Button onClick={handleSubmit} />
        </div>
      </div>
      <div className="">
        <Botbar/>
      </div>
    </>
  );
}

export default Login;
