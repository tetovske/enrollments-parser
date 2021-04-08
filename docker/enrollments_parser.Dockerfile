FROM default AS builder

WORKDIR /baumanfinder

CMD [".bin/enrollments-parser"]
