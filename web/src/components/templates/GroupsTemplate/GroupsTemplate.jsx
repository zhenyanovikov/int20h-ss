import { Box, Stack, CircularProgress } from "@mui/material";
import Empty from "../../molecules/Empty/Empty";
import CreateGroupForm from "../../organisms/CreateGroupForm/CreateGroupForm";
import AdminTemplate from "../../templates/AdminTemplate/AdminTemplate";

function GroupsTemplate({ groupsData, isGetGroupsLoading }) {
  return (
    <AdminTemplate>
      {isGetGroupsLoading || !groupsData ? (
        <Stack alignItems="center">
          <CircularProgress />
        </Stack>
      ) : (
        <Stack spacing={2}>
          <CreateGroupForm />

          <>{!groupsData.length ? <Empty /> : <Box />}</>
        </Stack>
      )}
    </AdminTemplate>
  );
}

export default GroupsTemplate;
