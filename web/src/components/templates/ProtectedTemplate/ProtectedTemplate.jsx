import { Navigate, Outlet } from "react-router-dom";
// import useIsLoggedIn from "../../../hooks/useIsLoggedIn";

function ProtectedTemplate() {
  // const isLoggedIn = useIsLoggedIn();
  const isLoggedIn = true;

  if (!isLoggedIn) {
    return <Navigate to="/" />;
  }

  return <Outlet />;
}

export default ProtectedTemplate;
