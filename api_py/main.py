from fastapi import FastAPI
from middleware import register_middleware
from exception import register_exception_handler
from router import register_router
from config import config
import uvicorn


from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker

from sqlalchemy import Boolean, Column, Integer, String, ForeignKey
from sqlalchemy.orm import relationship


print(config().database_default_url)
engine = create_engine(config().database_default_url)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)
Base = declarative_base()

class Test(Base):
    __tablename__ = "test"
    sceneId = Column(Integer, primary_key=True)
    sceneName = Column(String)
    sceneCode = Column(String, unique=True, index=True)
    sceneConfig = Column(String)
    isStop = Column(Boolean, default=False)
    rels = relationship("TestRel", back_populates="tests")

class TestRel(Base):
    __tablename__ = "test_rel"
    relId = Column(Integer, primary_key=True)
    relName = Column(String)
    sceneId = Column(Integer, ForeignKey("test.sceneId"))
    tests = relationship("Test", back_populates="rels")


app = FastAPI(docs_url=None, redoc_url="/redoc")
register_middleware(app)
register_exception_handler(app)
register_router(app)

# 启动方式（两种）
# 一般用于调试：uvicorn main:app --host=0.0.0.0 --port=20080 --reload
# 线上服务器用：python3.12 main.py
if __name__ == "__main__":
    if config().is_dev:
        uvicorn.run(app="main:app", host="0.0.0.0", port=8000, reload=True)
    else:
        uvicorn.run(app="main:app", host="0.0.0.0", port=20080)
