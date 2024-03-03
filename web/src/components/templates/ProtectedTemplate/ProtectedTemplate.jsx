import { Navigate, Outlet } from "react-router-dom";
import useIsAuthenticated from "../../../hooks/useIsAuthenticated";

function ProtectedTemplate() {
  const isAuthenticated = useIsAuthenticated();

  if (!isAuthenticated) {
    return <Navigate to="/" />;
  }

  return <Outlet />;
}

export default ProtectedTemplate;
