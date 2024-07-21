
// eslint-disable-next-line react/prop-types
function Button({onClick,title}) {


  return (
    <button type="button" onClick={onClick} className={`container  p-3 mb-1 border mt-3 border-slate-500 rounded-xl font-semibold hover:bg-sky-300`}>
    {title ? title : "Submit"}    
  </button>
)
}

export default Button