import { Card, CardContent, Typography, Link } from "@mui/material";
import { Link as RouterLink, generatePath } from "react-router-dom";
import { ROUTE } from "../../../constants/router";
import { useTranslation } from "react-i18next";

function Faculty({ group }) {
  const { t } = useTranslation();

  return (
    <Card>
      <CardContent>
        <Link
          variant="h6"
          component={RouterLink}
          underline="none"
          to={generatePath(ROUTE.GROUP_ANALYITCS, { id: group.id })}
          sx={{
            WebkitBoxOrient: "vertical",
            WebkitLineClamp: {
              xs: "2",
              md: "1",
            },
            display: "-webkit-box",
            overflow: "hidden",
          }}
        >
          {group.name}
        </Link>
        <Typography variant="body2" color="text.secondary">
          {group.faculty.name} • {group.yearOfStudy}{" "}
          {t("molecules.group.text.course")}
        </Typography>
      </CardContent>
    </Card>
  );
}

export default Faculty;
