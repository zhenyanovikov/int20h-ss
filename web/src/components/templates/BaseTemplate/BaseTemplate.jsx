import { Outlet } from "react-router-dom";
import Navbar from "../../organisms/Navbar/Navbar";

function BaseTemplate() {
  return (
    <>
      <Navbar />

      <Outlet />
    </>
  );
}

export default BaseTemplate;
