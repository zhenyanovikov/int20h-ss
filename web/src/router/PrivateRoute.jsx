import { Navigate } from "react-router-dom";
// import useIsLoggedIn from "../../../hooks/useIsLoggedIn";

function PrivateRoute({ children }) {
  // const isLoggedIn = useIsLoggedIn();
  const isLoggedIn = true;

  if (!isLoggedIn) {
    return <Navigate to="/" />;
  }

  return <>{children}</>;
}

export default PrivateRoute;
