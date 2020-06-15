FROM scratch
COPY http-sneak /http-sneak
ENTRYPOINT ["/http-sneak"]