import { BrowserRouter, Route, Routes } from "react-router-dom";
import BaseTemplate from "../components/templates/BaseTemplate/BaseTemplate";
import ProtectedTemplate from "../components/templates/ProtectedTemplate/ProtectedTemplate";
import EntryPage from "../components/pages/EntryPage/EntryPage";
import NotFoundPage from "../components/pages/NotFoundPage/NotFoundPage";
import FacultiesPage from "../components/pages/FacultiesPage/FacultiesPage";
import GroupsPage from "../components/pages/GroupsPage/GroupsPage";
import TeachersPage from "../components/pages/TeachersPage/TeachersPage";
import StudentsPage from "../components/pages/StudentsPage/StudentsPage";
import EmailTemplatesPage from "../components/pages/EmailTemplatesPage/EmailTemplatesPage";
import CreateEmailTemplateTemplate from "../components/templates/CreateEmailTemplateTemplate/CreateEmailTemplateTemplate";
import SendEmailPage from "../components/pages/SendEmailPage/SendEmailPage";
import { ROUTE } from "../constants/router";
import PrivateRoute from "./PrivateRoute";
import GroupAnalyticsTemplate from "../components/templates/GroupAnalyticsTemplate/GroupAnalyticsTemplate";

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
              path={ROUTE.GROUP_ANALYITCS}
              element={
                <PrivateRoute>
                  <GroupAnalyticsTemplate />
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
            <Route
              path={ROUTE.SEND_EMAIL}
              element={
                <PrivateRoute>
                  <SendEmailPage />
                </PrivateRoute>
              }
            />
            <Route
              path={ROUTE.EMAIL_TEMPLATES}
              element={
                <PrivateRoute>
                  <EmailTemplatesPage />
                </PrivateRoute>
              }
            />
            <Route
              path={ROUTE.CREATE_EMAIL_TEMPLATE}
              element={
                <PrivateRoute>
                  <CreateEmailTemplateTemplate />
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
