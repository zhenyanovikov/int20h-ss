import { Avatar, Stack, Typography } from "@mui/material";
import { getFullName } from "../../../helpers/user";

function UserPreview({ user }) {
  const fullName = getFullName(user.firstName, user.lastName);

  return (
    <Stack direction="row" alignItems="center" spacing={1}>
      <Avatar
        alt={fullName}
        src={user.avatarUrl}
        sx={{ width: 24, height: 24 }}
      />
      <Typography variant="body2">{fullName}</Typography>
    </Stack>
  );
}

export default UserPreview;
