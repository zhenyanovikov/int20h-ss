import { Box, Stack, CircularProgress } from "@mui/material";
import Empty from "../../molecules/Empty/Empty";
import InviteTeacherForm from "../../organisms/InviteTeacherForm/InviteTeacherForm";
import AdminTemplate from "../../templates/AdminTemplate/AdminTemplate";

function TeachersTemplate({ teachersData, isGetTeachersLoading }) {
  return (
    <AdminTemplate>
      {isGetTeachersLoading || !teachersData ? (
        <Stack alignItems="center">
          <CircularProgress />
        </Stack>
      ) : (
        <Stack spacing={2}>
          <InviteTeacherForm />

          <>{!teachersData.length ? <Empty /> : <Box />}</>
        </Stack>
      )}
    </AdminTemplate>
  );
}

export default TeachersTemplate;
