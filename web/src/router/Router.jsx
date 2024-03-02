import { BrowserRouter, Route, Routes } from "react-router-dom";
import BaseTemplate from "../components/templates/BaseTemplate/BaseTemplate";
import ProtectedTemplate from "../components/templates/ProtectedTemplate/ProtectedTemplate";
import EntryPage from "../components/pages/EntryPage/EntryPage";
import NotFoundPage from "../components/pages/NotFoundPage/NotFoundPage";
import FacultiesPage from "../components/pages/FacultiesPage/FacultiesPage";
import GroupsPage from "../components/pages/GroupsPage/GroupsPage";
import TeachersPage from "../components/pages/TeachersPage/TeachersPage";
import StudentsPage from "../components/pages/StudentsPage/StudentsPage";
import { ROUTE } from "../constants/router";
import PrivateRoute from "./PrivateRoute";

function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<BaseTemplate />}>
          <Route path="/" element={<EntryPage />} />
          <Route element={<ProtectedTemplate />}>
            <Route
              path={ROUTE.FACULTIES}
              element={
                <PrivateRoute>
                  <FacultiesPage />
                </PrivateRoute>
              }
            />
            <Route
              path={ROUTE.GROUPS}
              element={
                <PrivateRoute>
                  <GroupsPage />
                </PrivateRoute>
              }
            />
            <Route
              path={ROUTE.TEACHERS}
              element={
                <PrivateRoute>
                  <TeachersPage />
                </PrivateRoute>
              }
            />
            <Route
              path={ROUTE.STUDENTS}
              element={
                <PrivateRoute>
                  <StudentsPage />
                </PrivateRoute>
              }
            />
          </Route>
        </Route>
        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default Router;
