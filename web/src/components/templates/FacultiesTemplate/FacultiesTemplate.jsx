import { Box, Stack, CircularProgress } from "@mui/material";
import Empty from "../../molecules/Empty/Empty";
import CreateFacultyForm from "../../organisms/CreateFacultyForm/CreateFacultyForm";
import AdminTemplate from "../../templates/AdminTemplate/AdminTemplate";

function FacultiesTemplate({ facultiesData, isGetFacultiesLoading }) {
  return (
    <AdminTemplate>
      {isGetFacultiesLoading || !facultiesData ? (
        <Stack alignItems="center">
          <CircularProgress />
        </Stack>
      ) : (
        <Stack spacing={2}>
          <CreateFacultyForm />

          <>{!facultiesData.length ? <Empty /> : <Box />}</>
        </Stack>
      )}
    </AdminTemplate>
  );
}

export default FacultiesTemplate;
