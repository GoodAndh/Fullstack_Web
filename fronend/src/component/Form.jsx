import React from "react";

// eslint-disable-next-line react/prop-types
function Form({ inputID, title, Type, OnInputChange, validate }) {
  const BaseClass ="mt-2 px-3 py-2 shadow-md rounded-xl w-full block text-md hover:cursor-pointer";
  const InvalidClass ="border-2 border-red-800 outline-red-800 peer";
  const invalidMsgClass ="text-sm m-3 font-semibold invisible text-pink-800 peer-valid:visible";
  
  const inputChange = React.useCallback(
    (e) => {
      if (Type === "file") {
        const { id, value, files } = e.target;
      
        OnInputChange(id, value, files);

      }
      const { id, value } = e.target;
      OnInputChange(id, value);
    },
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [OnInputChange]
  );

  return (
    <>
      {/* ⬇️⬇️ pakai ini untuk membungkus div ⬇️⬇️*/}
      {/* <div className="max-w-lg m-3 mx-auto p-5 border-2 border-slate-600 rounded-xl shadow-md ">
      </div> */}

      <label htmlFor={inputID}>
        <span className="m-3 font-semibold block hover:cursor-pointer after:content-['*'] after:text-pink-500">
          {title}
        </span>
        <input
          onChange={inputChange}
          id={inputID}
          type={Type ? Type : "text"}
          className={Type!=="file" && `${BaseClass} ${validate !== "" ? InvalidClass : ""} `}  
          // {...(Type!=="file" ? {className:notFileClassName}:{})}    
          {...(Type === "number" ? { min: "0" } : {})}
        />
        {/* jika false ada error ,otherwise true */}
        {validate !== "" && <span className={invalidMsgClass}>{validate}</span>}
      </label>
    </>
  );
}

export default Form;
