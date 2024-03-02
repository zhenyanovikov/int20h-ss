import { Box, Stack, CircularProgress } from "@mui/material";
import AdminTemplate from "../../templates/AdminTemplate/AdminTemplate";
import Empty from "../../molecules/Empty/Empty";

function GroupsTemplate({ groupsData, isGetGroupsLoading }) {
  return (
    <AdminTemplate>
      {isGetGroupsLoading || !groupsData ? (
        <Stack alignItems="center">
          <CircularProgress />
        </Stack>
      ) : (
        <>{!groupsData.length ? <Empty /> : <Box />}</>
      )}
    </AdminTemplate>
  );
}

export default GroupsTemplate;
