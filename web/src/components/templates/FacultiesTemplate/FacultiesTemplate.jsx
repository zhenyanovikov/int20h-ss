import { Stack, CircularProgress } from "@mui/material";
import Empty from "../../molecules/Empty/Empty";
import CreateFacultyForm from "../../organisms/CreateFacultyForm/CreateFacultyForm";
import AdminTemplate from "../../templates/AdminTemplate/AdminTemplate";
import Faculties from "../../organisms/Faculties/Faculties";

function FacultiesTemplate({ facultiesData, isGetFacultiesLoading }) {
  return (
    <AdminTemplate>
      {isGetFacultiesLoading || !facultiesData ? (
        <Stack alignItems="center">
          <CircularProgress />
        </Stack>
      ) : (
        <Stack spacing={4}>
          <CreateFacultyForm />

          <>
            {!facultiesData.length ? (
              <Empty />
            ) : (
              <Faculties faculties={facultiesData} />
            )}
          </>
        </Stack>
      )}
    </AdminTemplate>
  );
}

export default FacultiesTemplate;
