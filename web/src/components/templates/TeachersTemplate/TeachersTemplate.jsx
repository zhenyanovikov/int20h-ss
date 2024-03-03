import { Stack, CircularProgress } from "@mui/material";
import Empty from "../../molecules/Empty/Empty";
import InviteTeacherForm from "../../organisms/InviteTeacherForm/InviteTeacherForm";
import Teachers from "../../organisms/Teachers/Teachers";
import AdminTemplate from "../../templates/AdminTemplate/AdminTemplate";

function TeachersTemplate({ teachersData, isGetTeachersLoading }) {
  return (
    <AdminTemplate>
      {isGetTeachersLoading || !teachersData ? (
        <Stack alignItems="center">
          <CircularProgress />
        </Stack>
      ) : (
        <Stack spacing={4}>
          <InviteTeacherForm />

          <>
            {!teachersData.length ? (
              <Empty />
            ) : (
              <Teachers teachers={teachersData} />
            )}
          </>
        </Stack>
      )}
    </AdminTemplate>
  );
}

export default TeachersTemplate;
