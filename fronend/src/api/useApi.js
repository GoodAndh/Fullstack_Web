/* eslint-disable react-refresh/only-export-components */
import instance from "../utils/instance.js";



export const postApi = async (url,value) => {
  let response;
  let error;
  try {
    const re = await instance.post(url, value);
    response = re.data;
  } catch (e) {

    error = e.response.data;
  }

  return { error, response };
};

export const getApi = async (url,form=null,config={})=>{
  let response;
  let error;

  try {
    const re =await instance.get(url,{
      ...config,
      params:form
    })
    response=re.data;
  } catch (e) {
    error=e.response.data
  }

  return {response,error}
}