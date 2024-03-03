import { Stack } from "@mui/material";
import Group from "../../molecules/Group/Group";

function Groups({ groups }) {
  return (
    <Stack spacing={2}>
      {groups.map((group) => (
        <Group key={group.id} group={group} />
      ))}
    </Stack>
  );
}

export default Groups;
