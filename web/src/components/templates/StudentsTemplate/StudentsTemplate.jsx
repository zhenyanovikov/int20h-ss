import { Box, Stack, CircularProgress } from "@mui/material";
import AdminTemplate from "../../templates/AdminTemplate/AdminTemplate";
import Empty from "../../molecules/Empty/Empty";

function StudentsTemplate({ studentsData, isGetStudentsLoading }) {
  return (
    <AdminTemplate>
      {isGetStudentsLoading || !studentsData ? (
        <Stack alignItems="center">
          <CircularProgress />
        </Stack>
      ) : (
        <>{!studentsData.length ? <Empty /> : <Box />}</>
      )}
    </AdminTemplate>
  );
}

export default StudentsTemplate;
