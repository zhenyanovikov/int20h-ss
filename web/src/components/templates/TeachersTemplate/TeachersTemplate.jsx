import { Box, Stack, CircularProgress } from "@mui/material";
import AdminTemplate from "../../templates/AdminTemplate/AdminTemplate";
import Empty from "../../molecules/Empty/Empty";

function TeachersTemplate({ teachersData, isGetTeachersLoading }) {
  return (
    <AdminTemplate>
      {isGetTeachersLoading || !teachersData ? (
        <Stack alignItems="center">
          <CircularProgress />
        </Stack>
      ) : (
        <>{!teachersData.length ? <Empty /> : <Box />}</>
      )}
    </AdminTemplate>
  );
}

export default TeachersTemplate;
