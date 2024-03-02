import { Box, Stack, CircularProgress } from "@mui/material";
import AdminTemplate from "../../templates/AdminTemplate/AdminTemplate";
import InviteTeacherForm from "../../organisms/InviteTeacherForm/InviteTeacherForm";
import Empty from "../../molecules/Empty/Empty";

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
