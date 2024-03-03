import { Stack, CircularProgress } from "@mui/material";
import Empty from "../../molecules/Empty/Empty";
import CreateGroupForm from "../../organisms/CreateGroupForm/CreateGroupForm";
import Groups from "../../organisms/Groups/Groups";
import AdminTemplate from "../../templates/AdminTemplate/AdminTemplate";

function GroupsTemplate({ groupsData, isGetGroupsLoading }) {
  return (
    <AdminTemplate>
      {isGetGroupsLoading || !groupsData ? (
        <Stack alignItems="center">
          <CircularProgress />
        </Stack>
      ) : (
        <Stack spacing={4}>
          <CreateGroupForm />

          <>{!groupsData.length ? <Empty /> : <Groups groups={groupsData} />}</>
        </Stack>
      )}
    </AdminTemplate>
  );
}

export default GroupsTemplate;
