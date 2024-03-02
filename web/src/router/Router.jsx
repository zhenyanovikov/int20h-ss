import { BrowserRouter, Route, Routes } from "react-router-dom";
import EntryPage from "../components/pages/EntryPage/EntryPage";
import NotFoundPage from "../components/pages/NotFoundPage/NotFoundPage";

function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<EntryPage />} />
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default Router;
