/* eslint-disable no-prototype-builtins */
import React from "react";
import Navbar from "../component/Navbar.jsx";
import Form from "../component/Form.jsx";
import Button from "../component/Button.jsx";
import Botbar from "../component/Botbar.jsx";
import { postApi } from "../api/useApi.js";

function Register() {
  const [formValue, setFormValue] = React.useState({
    name: "",
    username: "",
    password: "",
    vpassword: "",
    email: "",
  });

  const [errorMsg, setError] = React.useState({
    errorname: "",
    errorusername: "",
    errorpassword: "",
    errorvpassword: "",
    erroremail: "",
  });

  const inputChange = (id, value) => {
    setFormValue((p) => ({
      ...p,
      [id]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault;

    const { error, response } = await postApi("register/users", formValue);
    if (error) {
      if (error.data) {
        Object.keys(error.data).forEach((keys) => {
          if (errorMsg.hasOwnProperty(keys)) {
            setError((p) => ({
              ...p,
              [keys]: error.data[keys],
            }));
          }
        });

        Object.keys(errorMsg).forEach((keys) => {
          if (!error.data.hasOwnProperty(keys)) {
            setError((p) => ({
              ...p,
              [keys]: "",
            }));
          }
        });
      }
    }

    if (response) {
      console.log("response:", response);
      setFormValue((p) => ({
        ...p,
        errorname: "",
        errorusername: "",
        errorpassword: "",
        errorvpassword: "",
        erroremail: "",
      }));
    }
  };
  return (
    <>
      <div className="sm:mb-20 md:mb:28 mt-1 md:mt-3">
        <Navbar />
      </div>
      <div className="max-w-lg m-3 mx-auto p-5 border-2 border-slate-600 rounded-xl shadow-md  ">
        <Form
          inputID={"name"}
          title={"Name"}
          Type={"text"}
          OnInputChange={inputChange}
          validate={errorMsg.errorname}
        />
        <Form
          inputID={"username"}
          title={"Username"}
          Type={"text"}
          OnInputChange={inputChange}
          validate={errorMsg.errorusername}
        />
        <Form
          inputID={"password"}
          title={"Password"}
          Type={"password"}
          OnInputChange={inputChange}
          validate={errorMsg.errorpassword}
        />
        <Form
          inputID={"vpassword"}
          title={"Confirm Password"}
          Type={"password"}
          OnInputChange={inputChange}
          validate={errorMsg.errorvpassword}
        />
        <Form
          inputID={"email"}
          title={"Email"}
          Type={"text"}
          OnInputChange={inputChange}
          validate={errorMsg.erroremail}
        />
        <div className="m-1">
          <Button onClick={handleSubmit} />
        </div>
      </div>
      <div className=" mt-36">
        <Botbar />
      </div>
    </>
  );
}

export default Register;
