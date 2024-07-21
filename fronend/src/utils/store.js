export const accessToken = localStorage.getItem("access_token");
export const refreshToken = localStorage.getItem("refresh_token");

export const saveToken = (key, value) => localStorage.setItem(key, value);

